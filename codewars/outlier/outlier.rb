def FindOutlier arr
  countOdd = 0
  countEven = 0

  arr.each do |num|
    if num.even?
      countEven += 1
    else
      countOdd += 1
    end
  end

  if countEven > countOdd
    arr.each do |num|
      return num if num.odd?
    end
  else
    arr.each do |num|
      return num if num.even?
    end
  end
end

puts FindOutlier([2, 6, 8, -10, 3])