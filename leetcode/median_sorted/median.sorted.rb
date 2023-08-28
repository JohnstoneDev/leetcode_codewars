def findMedianSortedArrs nums1, nums2
	merged = nums1.concat(nums2).sort
	median = 0

	if merged.length % 2 == 0
		x = merged.length / 2
		y = x - 1
		median = (merged[x] + merged[y]) / 2.0
	else
		middle_index = merged.length / 2
		median += merged[middle_index]
	end

	return median
end

puts "#{findMedianSortedArrs([1,3],[2])}"
puts "#{findMedianSortedArrs([1,2],[3,4])}"