// static types will ensure "character" is always a string
// types must stay same type
// types are inferred and do not need types in declaration
let character = "mario";
let isCharacterStr = true;

// "number" type handles floats & int
let age = 30;

// function type restrictions assigned in arguments "(arg: type[, arg: type ...])"
const circ = (diameter: number) => {
    return diameter * Math.PI;
};

// typescript has access to everything javascript does, like "Math", "console", & "document"
console.log(circ(10));

// arrays must only have the same types of data
let names = ["mario", "Luigi", "Peach"];
names.push("Yoshi");
//names.push(3); // will throw an error, because numbers cannot be added to str array
//names[1] = 100;

// you can infer multiple types in an array
let mixed = [1, "Ryu"];
//mixed.push(true); // only inferred types from declaration my be used

// objects can be inferred like arrays
let ninja = {
    name: "Ryu",
    belt: "black",
    age: 30
};
ninja.age = 40;