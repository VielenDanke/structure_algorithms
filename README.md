# Algorithms and Data Structures

## Base
````
log(n) < n < n^2
````

### Arrays
````
Insertion - in the end - O(1), other - O(n)
Removal - in the end - O(1), other - O(n)
Searching - O(n)
Access - O(1)
````

### Sorting
````
 Sort type               | Average time complexity (Best - Worst)           | Space complexity |
------------------------ | ------------------------------------------------ | ---------------- |
  Bubble                 | O(N^2)                  (O(N) - O(N^2))          | O(1)             |
  Insertion              | O(N^2)                  (O(N) - O(N^2))          | O(1)             |
  Selection              | O(N^2)                  (O(N^2) - O(N^2))        | O(1)             |
  Merge                  | O(N logN)               (O(N logN) - O(N logN))  | O(N)             |
  Quick                  | O(N logN)               (O(N logN) - O(N^2))     | O(logN)          |
  Radix (mostly numbers) | O(NK)                   (O(NK) - O(NK))          | O(N + K)         |  
````