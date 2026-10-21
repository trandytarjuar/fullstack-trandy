using System;
using System.Collections.Generic;

// Algorithm exercise: each method is independent and handles common edge cases.
public class StringReverser
{
    public static string ReverseString(string inputStr)
    {
        ArgumentNullException.ThrowIfNull(inputStr);
        // Reverse UTF-16 code units as commonly expected in introductory string exercises.
        // Complex Unicode grapheme clusters are not preserved by this simple approach.
        char[] chars = inputStr.ToCharArray();
        Array.Reverse(chars);
        return new string(chars);
    }
}

public class PalindromeChecker
{
    public static bool IsPalindrome(string inputStr)
    {
        ArgumentNullException.ThrowIfNull(inputStr);
        // Exact, case-sensitive comparison; whitespace and punctuation count.
        // Two pointers avoid allocating a reversed string: O(n) time, O(1) space.
        int left = 0, right = inputStr.Length - 1;
        while (left < right)
        {
            if (inputStr[left] != inputStr[right]) return false;
            left++;
            right--;
        }
        return true;
    }
}

public class PrimeNumberGenerator
{
    public static List<int> GeneratePrimes(int limit)
    {
        var primesList = new List<int>();
        if (limit < 2) return primesList;
        // Sieve of Eratosthenes: O(n log log n) time, O(n) memory.
        // Start crossing out at p*p; smaller multiples have already been crossed out.
        var composite = new bool[limit + 1];
        for (int p = 2; p <= limit / p; p++)
        {
            if (composite[p]) continue;
            for (int multiple = p * p; multiple <= limit; multiple += p)
                composite[multiple] = true;
        }
        for (int n = 2; n <= limit; n++)
            if (!composite[n]) primesList.Add(n);
        return primesList;
    }
}
