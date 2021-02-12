function factorialize(num) {
  if (num < 0) 
        return -1;
  else if (num == 0) 
      return 1;
  else {
      return (num * factorialize(num - 1));
  }
}

//console.log(factorialize(5))
function factorial(){
  var num ;
    num = document.getElementById("num").value;
    
  var cal ,text3 ;
    cal = document.getElementById("cal").value;
    cal = num * factorialize(num - 1) 
    text3 = cal;
    document.getElementById("demo3").innerHTML =  Math.round(text3);
}