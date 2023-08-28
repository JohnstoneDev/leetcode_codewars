function findMedianSortedArrs(arr1, arr2){
	let mergedArray = arr1.concat(arr2) // merge the two arrays
	mergedArray = mergedArray.sort() // sort the merger
	let median

	if (mergedArray.length % 2 == 0) {
		let x = mergedArray.length / 2
		let y = x - 1
		median = (mergedArray[x] + mergedArray[y]) / 2
	} else {
		let middleIndex = Math.floor(mergedArray.length / 2)
		median = mergedArray[middleIndex]
	}

	return median
}

console.log(findMedianSortedArrs([1,2],[3,4]))