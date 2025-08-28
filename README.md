# Challenge

Find all 3-digit numbers that satisfy two conditions:

1. When you rearrange its digits to form a new 3-digit number (a permutation), the new number must be larger than the original one.

2. The difference between the new number and the original number must also be a permutation of the original number's digits.

# Human language

John went to restaurant. In his pocket some amount of money, he ordered sushi (price of which is permutation of pocket money). How much money John came into restaurant with, if after dinner change is also permutation of the original amount?

# Negative example:
Let's use the number 145 to illustrate what's being asked.

**Original Number:** 145

**Permutations:** You can rearrange the digits (1, 4, 5) to create numbers like 154, 415, 451, 514, and 541.

**Condition 1:** The rearranged number must be larger. Let's use 451.

**Condition 2:** The difference between 451 and 145 must be a permutation of 1, 4, and 5.

451−145=306

The digits of 306 (3, 0, 6) are not the same as the digits of the original number (1, 4, 5).

So, 145 does **not** meet the criteria. The challenge is to find the numbers that do work.

# Positive example

321 - 123 = 231 (though math doesn't work out in this example)
