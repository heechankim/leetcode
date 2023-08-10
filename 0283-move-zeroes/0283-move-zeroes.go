func moveZeroes(nums []int)  {
  
  if len(nums) == 1 {
    return
  }
  
  zp := 0

  np := 1

  

  for {

    if nums[zp] == 0 {

      if nums[np] != 0 {

        nums[zp], nums[np] = nums[np], nums[zp]

        zp++

      } else {

        np++

      }

    } else {

      zp++

      np++

    }

    if zp == len(nums) || np == len(nums) {

      break

    }

    

  }
}