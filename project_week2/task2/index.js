export default function ReverseWord(str) {
    let result = ""
    let wordArray = str.split(" ")
    for(let i = wordArray.length-1; i >= 0; i--) {
        result += wordArray[i] + " "
    }
    return console.log(result)
}