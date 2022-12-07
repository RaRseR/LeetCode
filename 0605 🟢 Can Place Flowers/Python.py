def canPlaceFlowers(self, flowerbed: List[int], n: int) -> bool:
  c: int = 0
  f: int = len(flowerbed)
        
  for i in range(f):
    if flowerbed[i] == 0:
      l: bool = (i == 0) or (flowerbed[i - 1] == 0)
      r: bool = (i == f - 1) or (flowerbed[i + 1] == 0)

      if l and r:
        flowerbed[i] = 1
         c += 1
          
  return c >= n
