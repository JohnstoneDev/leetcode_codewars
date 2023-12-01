def ToJadenCase string
  words = string.split.map(&:capitalize).join(' ')
end


puts ToJadenCase("most trees are blue")
puts ToJadenCase("All the rules in this world were made by someone no smarter than you. So make your own.")
puts ToJadenCase("When I die. then you will realize")
puts ToJadenCase("Jonah Hill is a genius")
puts ToJadenCase("Dying is mainstream")

