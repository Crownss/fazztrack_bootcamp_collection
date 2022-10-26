export default function FazzFood(price, voucher, how_far, fee){
    if(typeof price != 'number' &&typeof  voucher != 'string' &&typeof  how_far != 'number' &&typeof  fee != 'boolean'){
        // console.log(typeof price, typeof voucher, typeof how_far, typeof fee)
        return console.log("usage FazzFood(75000, 'FAZZFOOD50', 5, true)")
    }
    let deliveryFee = 5000
    let discountA = price * 0.5
    let discountB = price * 0.6
    let limitDiscountA = price - 50000
    let limitDiscountB = price - 30000
    let feeTrue = price * 0.05

    if(how_far > 2){
        deliveryFee=((how_far-2)*3000)+deliveryFee
    }
    //a
    if(voucher.toUpperCase() == "FAZZFOOD50"){
        if(price >= 50000){
            if(discountA < 50000){
                if(fee){
                    return console.log(
                        "Harga: ",price+'\n',
                        "Potongan: ",discountA+'\n',
                        "Biaya Antar: ",deliveryFee+'\n',
                        "Pajak: ",feeTrue+'\n',
                        "SubTotal: ",deliveryFee+discountA+feeTrue+'\n'
                    )
                }else{
                    return console.log(
                        "Harga: ",price+'\n',
                        "Potongan: ",discountA+'\n',
                        "Biaya Antar: ",deliveryFee+'\n',
                        "Pajak: ",0+'\n',
                        "SubTotal: ",deliveryFee+discountA+'\n'
                    )
                }
            }else{
                if(fee){
                    return console.log(
                        "Harga: ",price+'\n',
                        "Potongan: ",50000+'\n',
                        "Biaya Antar: ",deliveryFee+'\n',
                        "Pajak: ",feeTrue+'\n',
                        "SubTotal: ",deliveryFee+limitDiscountA+feeTrue+'\n'
                    )
                }else{
                    return console.log(
                        "Harga: ",price+'\n',
                        "Potongan: ",50000+'\n',
                        "Biaya Antar: ",deliveryFee+'\n',
                        "Pajak: ",0+'\n',
                        "SubTotal: ",deliveryFee+limitDiscountA+'\n'
                    )
                }
            }

        }else{
            if(fee){
                return console.log(
                    "Harga: ",price+'\n',
                    "Potongan: ",0+'\n',
                    "Biaya Antar: ",deliveryFee+'\n',
                    "Pajak: ",feeTrue+'\n',
                    "SubTotal: ",deliveryFee+price+feeTrue+'\n'
                )
            }else{
                return console.log(
                    "Harga: ",price+'\n',
                    "Potongan: ",0+'\n',
                    "Biaya Antar: ",deliveryFee+'\n',
                    "Pajak: ",0+'\n',
                    "SubTotal: ",deliveryFee+price+'\n'
                )
            }
        }
    }
    //b
    else if(voucher.toUpperCase() == "DITRAKTIR60"){
        if(price >= 25000){
            if(discountB < 30000){
                if(fee){
                    return console.log(
                        "Harga: ",price+'\n',
                        "Potongan: ",discountB+'\n',
                        "Biaya Antar: ",deliveryFee+'\n',
                        "Pajak: ",feeTrue+'\n',
                        "SubTotal: ",deliveryFee+discountB+feeTrue+'\n'
                    )
                }else{
                    return console.log(
                        "Harga: ",price+'\n',
                        "Potongan: ",discountB+'\n',
                        "Biaya Antar: ",deliveryFee+'\n',
                        "Pajak: ",0+'\n',
                        "SubTotal: ",deliveryFee+discountB+'\n'
                    )
                }
            }else{
                if(fee){
                    return console.log(
                        "Harga: ",price+'\n',
                        "Potongan: ",30000+'\n',
                        "Biaya Antar: ",deliveryFee+'\n',
                        "Pajak: ",feeTrue+'\n',
                        "SubTotal: ",deliveryFee+limitDiscountB+feeTrue+'\n'
                    )
                }else{
                    return console.log(
                        "Harga: ",price+'\n',
                        "Potongan: ",30000+'\n',
                        "Biaya Antar: ",deliveryFee+'\n',
                        "Pajak: ",0+'\n',
                        "SubTotal: ",deliveryFee+limitDiscountB+'\n'
                    )
                }
            }
        }else{
            if(fee){
                return console.log(
                    "Harga: ",price+'\n',
                    "Potongan: ",0+'\n',
                    "Biaya Antar: ",deliveryFee+'\n',
                    "Pajak: ",feeTrue+'\n',
                    "SubTotal: ",deliveryFee+price+feeTrue+'\n'
                )
            }else{
                return console.log(
                    "Harga: ",price+'\n',
                    "Potongan: ",0+'\n',
                    "Biaya Antar: ",deliveryFee+'\n',
                    "Pajak: ",0+'\n',
                    "SubTotal: ",deliveryFee+price+'\n'
                )
            }
        }
    }else{
        if(fee){
            return console.log(
                "Harga: ",price+'\n',
                "Potongan: ",0+'\n',
                "Biaya Antar: ",deliveryFee+'\n',
                "Pajak: ",feeTrue+'\n',
                "SubTotal: ",deliveryFee+price+feeTrue+'\n'
            )
        }else{
            return console.log(
                "Harga: ",price+'\n',
                "Potongan: ",0+'\n',
                "Biaya Antar: ",deliveryFee+'\n',
                "Pajak: ",0+'\n',
                "SubTotal: ",deliveryFee+price+'\n'
            )
        }
    }
}