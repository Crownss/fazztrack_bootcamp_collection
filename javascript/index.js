w//Task 1
const biodata = {
    name: "Djezen RHD",
    age: 19,
    hobbies: ["berenang", "ngegame"],
    schoolList: [
        {
            name: "SDN 03 TANAH TINGGI",
            yearIn: 2010,
            yearOut: 2016,
            major: null
        },
        {
            name: "SMPN 2 JAKARTA PUSAT",
            yearnIn: 2016,
            yearOut: 2019,
            major: null
        },
        {
            name: "SMKN 2 JAKARTA PUSAT",
            yearIn: 2019,
            yearOut: 2022,
            major: "RPL (Rekayasa Perangkat Lunak)"
        }
    ],
    skills: [
        {
            skillName: "Python",
            level: "expert"
        },
        {
            skillName: "Javascript",
            level: "advance"
        },
        {
            skillName: "Golang",
            level: "advance"
        }
    ]
}

//Task 2
const mtk = 80
const bhsaIndonesia = 85
const bhsaInggris = 75
const ipa = 80
if (typeof mtk != Number && typeof bhsaIndonesia != Number && typeof bhsaInggris != Number && typeof ipa != Number )return console.log("data harus number")
let mean = (mtk + bhsaIndonesia + bhsaInggris + ipa) /4
console.log("nilai rata - rata: "+mean)
if (mean >= 90 && mean <= 100) {
    console.log("grade: A")
}else if(mean >= 80 && mean <= 89){
    console.log("grade: B")
}else if(mean >= 70 && mean <= 79){
    console.log("grade: C")
}else if(mean >= 60 && mean <= 69){
    console.log("grade: D")
}else{
    console.log("grade: E")
}

//Task 3
function reversalTriangle(n){
    if (typeof n != Number)return console.log("error harus number")
    let a = '';
    for(let i = n; i >= 1; i--){
        for(let j = 0; j < i; j++){
          a += i
        }
        console.log(a);
        a = '';
    }
  }
reversalTriangle("enam")

//Task 4
//a.
let data = {
    id: 1,
    name: "Leanne Graham",
    username: "Bret",
    email: "Sincere@april.biz",
    hobby: ["ngegame", "berenang"],
    address:{
        street: "Kulas Light",
        suite: "Apt. 556",
        city: "Gwenborough",
        zipcode: "92998-3874",
    },
    phone: "1-770-736-8031 x56442",
    website: "hildegard.org",
}
const {...datas}= data
console.log(...datas.name, ...datas.email, ...datas.hobby)

//b
let street = data.address.street
let city = data.address.city
console.log(street, city)