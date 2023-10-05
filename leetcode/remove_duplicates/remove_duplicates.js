
function removeDuplicates(numbers){
  if (numbers.length == 0){
    return 0
  }

  let k = 1

  for (let i = 1; i < numbers.length; i++ ){
    if (numbers[i] != numbers[i -1]){
      numbers[k] = numbers[i]
      k++
    }
  }

  return k
}


console.log(removeDuplicates([1,1,2]))
console.log(removeDuplicates([0, 0, 1, 1, 1, 2, 2, 3, 3, 4]))