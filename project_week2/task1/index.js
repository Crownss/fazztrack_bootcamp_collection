export default function CheckPalindrome(str){
    if (str === ''){
        return console.log('string tidak boleh kosong')
    }else if(str.length === 1 ){
        return console.log('string harus lebih dari 1')
    }
    let temp = []
    let belakang = ''
    let depan = ''
    let indikator = true

    for (let i = 0; i < str.length; i++) {
        temp.push(str[i].toUpperCase())
    }
    while(temp.length > 1 && indikator){
        depan += temp.shift()
        belakang += temp.pop()
        if(depan != belakang){
            indikator = false
        }
    }
    if(indikator){
        return console.log(`${str} is palindrom`)
    }else{
        return console.log(`${str} is not palindrom`)
    }
}