# Leetcode problems README

This is basically all of the leetcode execises I have ever done (in Go).


## Takeaways

- Numbers: use smaller `int`s, `uint`s, and `float`s whenever possible. It has suprisingly high impact on performance.
- Counting characters: using arrays is more computationally and space efficient than hashmaps.  
  ```go
  counts := make([]int, 26)
  counts[char - 'a']++
  ```
- Bit masking:
  - XOR two numbers to find their different bits. `A ^ B`. If it's 0, the numbers are identical.  
    Very useful for finding / ignoring pairs.
- Dynamic programming: it's about incrementally solving a problem by first
  solving a much simpler subset of your problem. "Divide and conquer".
  Sometimes it's obvious, sometimes it makes no sense to me.  
  It's not going to work on its own with huge amounts of possibilities. You
  want to look at memoization solutions or other solutions (other than dynamic
  programming) entirely.
- Memoization: are you doing a bunch of calculations that might reoccur several
  times while solving your problem? Just create a structure to store these
  intermediate results, and reuse them instead of doing the calculation again.
- Priority queues: track / prioritise past decisions and possibly change your mind about them.


Sidenotes:
- The go documentation is mediocre.
- The heap package was designed to be confusing. Garbage user experience.
