function toCalculatorBmi()
{
    var weight = prompt("น้ำหนัก : ") 
    var height = prompt("ส่วนสูง : ")
    let value =   weight / (height /100 ) ** 2
    return value.toFixed(2) ;
}
//console.log(toCalculatorBmi())

function myFunction() {
    var weight, text ;
    weight = document.getElementById("weight").value;
    //text = weight;
    //document.getElementById("demo").innerHTML = text;

    var height,text2 ; 
    height = document.getElementById("height").value;
    //text2 = height;
    //document.getElementById("demo2").innerHTML = text2;

    var cal ,text3 ;
    cal = document.getElementById("cal").value;
    cal = weight / (height /100 ) ** 2
    text3 = cal;
    document.getElementById("demo3").innerHTML =  Math.round(text3);
    
} 



    
