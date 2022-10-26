//Task 1
let names = "djezen ramsey hamar della"
let nameInArr = [10,231,32,1,123,67,98,86]
/*
1.length menghitung panjang dari karakter -> names.length
2.split memecahkan karakter menjadi array -> names.split(" ")
3.toUpperCase mengubah karakter menjadi huruf kapital semua -> names.toUpperCase()
4.toLowerCase mengubah karakter menjadi huruf kecil semua -> names.toLowerCase()
5.sort mengurutkan array -> nameInArr.sort()
6.reverse memutar balikan array -> nameInArr.reverse()
7.join menambah karakter atau apapun disetiap array -> nameInArr.join(" + ")
8.splice menambahkan/menghapus isi array -> nameInArr.splice()
9.shift menghapus dan mengambil karakter pertama dari array -> nameInArr.shift()
10.pop menghapus dan mengambil karakter terakhir dari array -> nameInArr.pop()
*/
console.log("\t\tTask 1\n","\n=======================================\n"+"1.length\n\tcontoh: name.length hasil ->",names.length,"\n\n2.split\n\tcontoh: name.split() hasil ->",names.split(" "),"\n\n3.toUpperCase\n\tcontoh: name.toUpperCase() hasil ->",names.toUpperCase(),"\n\n4.toLowerCase\n\tcontoh: name.toLowerCase() hasil ->",names.toLowerCase(),"\n\n5.sort\n\tcontoh: nameInArr.sort() hasil ->",nameInArr.sort(),"\n\n6.reverse\n\tcontoh: nameInArr.reverse() hasil ->",nameInArr.reverse(),"\n\n7.join\n\tcontoh: nameInArr.join(' + ') hasil ->",nameInArr.join(" + "),"\n\n8.splice\n\tcontoh: nameInArr.splice(2,0, 'sepuluh ribu') hasil ->",nameInArr.splice(2,0,"sepuluh ribu"),"\n\n9.pop\n\tcontoh: nameInArr.pop() hasil ->",nameInArr.pop(),"\n\n10.shift\n\tcontoh: nameInArr.shift() hasil ->",nameInArr.shift()+"\n=======================================\n\n")
//Task 2
let peopleName = ['Abigail', 'Alexandra', 'Alison','Amanda', 'Angela', 'Bella','Carol', 'Caroline', 'Carolyn','Deirdre', 'Diana', 'Elizabeth','Ella', 'Faith', 'Olivia', 'Penelope']
function cbFn(...params){
    console.log("\t\tTask 2","\n=======================================\n"+params+"\n=======================================\n\n")
}
function searchName(name, howmuch, callbackFn){
    if(name.length < 2)return console.log("huruf yg dicari harus lebih dari 2")
    if(howmuch < 1)return console.log("nama yg dicari harus lebih dari 1")
    let temp = []
    let sorting = peopleName.sort()
    let result = sorting.filter((value)=>{return value.match(name)})
    for(let i=0;i<=howmuch-1;i++){
        if(result[i]!=undefined){
            return temp.push(result[i])
        }else if(result.length == 0){
            return console.log("\t\tTask 2","\n=======================================\ndata not found !\n=======================================\n\n")
        }
    }return callbackFn(temp)
}
searchName("xs", 6, cbFn)

//Task 3
let list = [23,12,10,32,45,65,57,98,75,80]
function seleksiNilai(nilaiAwal, nilaiAkhir, arr){
    if(isNaN(nilaiAwal)&&isNaN(nilaiAkhir))return console.log("nilai awal dan akhir harus berupa angka")
    if(nilaiAkhir < nilaiAwal && nilaiAwal > nilaiAkhir)return console.log("nilai akhir harus lebih besar dari nilai awal")
    if(arr.length < 5)return console.log("jumlah data dalam array harus lebih 5")
    arr.sort()
    let result = arr.filter((value) =>{return value >= nilaiAwal && value <= nilaiAkhir})

    return console.log("\t\tTask 3","\n=======================================\n"+result+"\n=======================================\n\n")
}
seleksiNilai(40,60,list)