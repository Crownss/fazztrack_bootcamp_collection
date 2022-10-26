//Task 1
const cekHariKerja = (day) =>{return new Promise((resolve, reject) => {setTimeout(()=>{const dataDay = ['senin', 'selasa','rabu','kamis','jumat'];let cek = dataDay.find((item)=>{return item === day});if(cek)return resolve(cek);else return reject(new Error('hari ini bukan hari kerja'))},3000)})}

let result = (date) =>{cekHariKerja(date).then((res)=>{return console.log(res)}).catch((res)=>{return console.log(res.message)})}
let resultTry = async (date) =>{try{let cek = await cekHariKerja(date);return console.log(cek)}catch(err){return console.log(err.message)}}
resultTry('rabu')
result('jumat')


//Task 2
const getMonth = (callback)=>{setTimeout(()=>{let error = false;let month = ['januari','februari', 'maret','april','mei','juni','juli','agustus','september','oktober','november','desember'];if(!error)return callback(null, month);else return callback(new Error('sorry data not found'), [])},4000)}

const  callbackFn = (err, arr)=>{if(!err)return arr.map((res)=>{return console.log(res)});else return console.log(err.message)}
getMonth(callbackFn)

//Task 3
//a
const cekHariLibur = (day) =>{
    return new Promise((resolve, reject) => {
        setTimeout(()=>{
            const dataDay = ['sabtu', 'minggu']
            let cek = dataDay.find((item)=>{
                return item === day
            })
            if(cek)return resolve(cek)
            else return reject(new Error('hari ini bukan hari libur'))
        },5000)
    })
}
let task3_a = (date) =>{ cekHariLibur(date).then((res)=>{return console.log(res)}).catch((res)=>{return console.log(res.message)})}
task3_a('minggu')
//b
const trainerCewe = (trainerNameCewe) =>{
    return new Promise((resolve, reject) => {
        setTimeout(()=>{
            const namaTrainerCewe = ['ka winda', 'ka faela', 'ka putri']
            let cek = namaTrainerCewe.find((item)=>{
                return item === trainerNameCewe
            })
            if(cek)return resolve(cek)
            else return reject(new Error('itu bukan nama cewe'))
        },6000)
    })
}
let task3_b = (trainerName) =>{trainerCewe(trainerName).then((res)=>{return console.log(res)}).catch((res)=>{return console.log(res.message)})}
task3_b('ka putri')

//Task 4
const restApi = async (link) =>{const resultLink = await fetch(link);let data = await resultLink.json();let mapping = await data ? data.map((res)=>{return console.log(res.name)}):console.log("data not found !");return await mapping}
restApi('https://jsonplaceholder.typicode.com/users')