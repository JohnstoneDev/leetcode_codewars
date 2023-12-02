function FindOutlier(array) {
  let countOdd = 0
  let countEven = 0
  let value = 0

  array.map(num => {
    if(num % 2 == 0) {
      countEven++
    } else {
      countOdd++
    }
  })

  if (countEven > countOdd) {
    array.map((num) => {
      if (num % 2 !== 0) {
        value += num
      }
    })
  } else {
    array.map((num) => {
      if (num % 2 == 0) {
        value += num
      }
    })
  }

  return value
}


console.log(FindOutlier([2, 6, 8, -10, 3]))
