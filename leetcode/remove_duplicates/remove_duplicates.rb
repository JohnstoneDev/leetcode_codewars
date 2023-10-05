def remove_duplicates nums
  return 0 if nums.length == 0

  k = 0

  nums.each_with_index do |num, index|
    if index == 0 || num != nums[index - 1]
      nums[k] = num
      k += 1
    end
  end

  k
end

puts "#{remove_duplicates([1,1,2])}"
puts "#{remove_duplicates([0, 0, 1, 1, 1, 2, 2, 3, 3, 4])}"
