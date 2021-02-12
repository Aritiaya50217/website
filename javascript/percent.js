function display(elementID,value)
{
    document.getElementById(elementID).innerHTML = value
}
function Calculate(p,c)
{
    var p = prompt("ราคาเต็มของสินค้า : ")
    var c = prompt("เปอร์เซนต์ส่วนลด : ")

    let value = p * (c / 100)
    return  "ส่วนลด "+ value.toFixed(2) + " THB" ;
}

//console.log(Calculate())

function myFunction() {
    var price, text ;
    price = document.getElementById("price").value;

    var percent,text2 ; 
    percent = document.getElementById("percent").value;
    
    var cal ,text3 ;
    cal = document.getElementById("cal").value;
    cal = price * (percent / 100)
    text3 = cal;
    document.getElementById("demo3").innerHTML =  Math.round(text3) + " THB";
    
} 