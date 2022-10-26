export default function DevideAndSort(num){
    if(typeof num != 'number')return console.log("harus angka !")
    let SplitNum = num.toString().split(0)
    for (let i = 0; i < SplitNum.length; i++) {
        SplitNum[i] = SplitNum[i].split("").sort().join("")
    }
    console.log(parseInt(SplitNum.join("")))
}