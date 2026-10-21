Console.WriteLine("=== Algorithm Demo ===");

Console.Write("Enter text to reverse: ");
string textToReverse = Console.ReadLine() ?? string.Empty;
Console.WriteLine($"Reverse \"{textToReverse}\": {StringReverser.ReverseString(textToReverse)}");

Console.Write("Enter text to check for palindrome: ");
string palindromeText = Console.ReadLine() ?? string.Empty;
Console.WriteLine($"Is \"{palindromeText}\" a palindrome? {PalindromeChecker.IsPalindrome(palindromeText)}");

int primeLimit;
while (true)
{
    Console.Write("Enter the prime number limit: ");
    if (int.TryParse(Console.ReadLine(), out primeLimit) && primeLimit >= 0)
        break;

    Console.WriteLine("Please enter a non-negative integer.");
}

var primes = PrimeNumberGenerator.GeneratePrimes(primeLimit);
Console.WriteLine($"Prime numbers up to {primeLimit}: {string.Join(", ", primes)}");
