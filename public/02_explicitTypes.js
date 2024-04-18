"use strict";
// types can be declared "explicitly"
let character;
let age;
let isLoggedIn;
// declared "variables" can be initialized with a value
character = "Mario;";
// "arrays" may also be given an explicit type
let ninjas;
// ninjas.push("Luigi"); // cannot push to un initialized array
// declared variables can also be "initialized"
let names = ["Mario", "Luigi"];
names.push("Toad");
// "unions" allow multiple types
let mixedArray = ["Mario"]; // mixed types array declaration
mixedArray.push(20);
let mixedVariable = 12; // no parenthesis required for non array
mixedVariable = "12";
// "objects" may be explicitly typed as object
let ninjaOne;
ninjaOne = []; // arrays are also objects
ninjaOne = { name: "mario", age: 25 };
// objects may also explicitly have all parts defined
let NinjaTwo;
NinjaTwo = { name: "Luigi", age: 25, belt: true };
// "any" types are fully dynamic
let dynamicType = "string";
dynamicType = true;
dynamicType = 25;
dynamicType = ["don't", "use", "this", "type", "unless", "necessary"];
// "any" may also be used to arrays and objects
let dynamicArray = ["name", 50, true];
let ninjaThree;
ninjaThree = { name: 50, age: "yoshi" };
