function ToJadenCase(string) {
  const words = string.split(" ").map(word => {
    let letters = word.split("")

    letters[0] = letters[0].toUpperCase()

    return letters.join("")
  }).join(" ");

  return words
}


console.log(ToJadenCase("most trees are blue"))
console.log(ToJadenCase("All the rules in this world were made by someone no smarter than you. So make your own."))
console.log(ToJadenCase("When I die. then you will realize"))
console.log(ToJadenCase("Jonah Hill is a genius"))
console.log(ToJadenCase("Dying is mainstream"))