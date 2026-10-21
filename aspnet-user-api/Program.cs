using System.ComponentModel.DataAnnotations;
using Microsoft.EntityFrameworkCore;

var builder = WebApplication.CreateBuilder(args);
builder.Services.AddDbContext<UserDb>(options =>
    options.UseSqlite(builder.Configuration.GetConnectionString("Users") ?? "Data Source=users.db"));
var app = builder.Build();

// A small take-home application: ensure the SQLite schema exists at startup.
// Production systems should use versioned EF Core migrations instead.
using (var scope = app.Services.CreateScope())
    scope.ServiceProvider.GetRequiredService<UserDb>().Database.EnsureCreated();

var users = app.MapGroup("/api/users");
users.MapGet("/", async (UserDb db) =>
    Results.Ok(await db.Users.AsNoTracking().OrderBy(u => u.UserId).ToListAsync()));

users.MapGet("/{userId:int}", async (int userId, UserDb db) =>
    await db.Users.AsNoTracking().FirstOrDefaultAsync(u => u.UserId == userId) is { } user
        ? Results.Ok(user) : Results.NotFound(new { message = "User not found" }));

users.MapPost("/", async (UserRequest request, UserDb db) =>
{
    var errors = Validate(request);
    if (errors.Count > 0) return Results.ValidationProblem(errors);
    string email = request.Email.Trim().ToLowerInvariant();
    if (await db.Users.AnyAsync(u => u.Email == email))
        return Results.Conflict(new { message = "Email already exists" });
    var user = new User { FirstName = request.FirstName.Trim(), LastName = request.LastName.Trim(), Email = email };
    db.Users.Add(user);
    await db.SaveChangesAsync();
    return Results.Created($"/api/users/{user.UserId}", user);
});

users.MapPut("/{userId:int}", async (int userId, UserRequest request, UserDb db) =>
{
    var errors = Validate(request);
    if (errors.Count > 0) return Results.ValidationProblem(errors);
    var user = await db.Users.FindAsync(userId);
    if (user is null) return Results.NotFound(new { message = "User not found" });
    string email = request.Email.Trim().ToLowerInvariant();
    if (await db.Users.AnyAsync(u => u.UserId != userId && u.Email == email))
        return Results.Conflict(new { message = "Email already exists" });
    user.FirstName = request.FirstName.Trim();
    user.LastName = request.LastName.Trim();
    user.Email = email;
    await db.SaveChangesAsync();
    return Results.Ok(user);
});

users.MapDelete("/{userId:int}", async (int userId, UserDb db) =>
{
    var user = await db.Users.FindAsync(userId);
    if (user is null) return Results.NotFound(new { message = "User not found" });
    db.Users.Remove(user);
    await db.SaveChangesAsync();
    return Results.NoContent();
});

app.Run();

static Dictionary<string, string[]> Validate(UserRequest request)
{
    var errors = new Dictionary<string, string[]>();
    if (string.IsNullOrWhiteSpace(request.FirstName)) errors["firstName"] = ["First name is required"];
    if (string.IsNullOrWhiteSpace(request.LastName)) errors["lastName"] = ["Last name is required"];
    if (string.IsNullOrWhiteSpace(request.Email) ||
        !new EmailAddressAttribute().IsValid(request.Email.Trim()))
        errors["email"] = ["A valid email is required"];
    if (request.FirstName?.Length > 100) errors["firstName"] = ["Maximum 100 characters"];
    if (request.LastName?.Length > 100) errors["lastName"] = ["Maximum 100 characters"];
    if (request.Email?.Length > 254) errors["email"] = ["Maximum 254 characters"];
    return errors;
}

public sealed record UserRequest(string? FirstName, string? LastName, string? Email);
public sealed class User
{
    public int UserId { get; set; }
    public string FirstName { get; set; } = "";
    public string LastName { get; set; } = "";
    public string Email { get; set; } = "";
}
public sealed class UserDb(DbContextOptions<UserDb> options) : DbContext(options)
{
    public DbSet<User> Users => Set<User>();
    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.Entity<User>().HasKey(u => u.UserId);
        modelBuilder.Entity<User>().HasIndex(u => u.Email).IsUnique();
    }
}
