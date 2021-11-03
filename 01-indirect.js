var foodSelections ={
    BerriesAndThings: ['orange','peach','tomato','potato','kiwi','strawberry','ground beer','blackberry','cherry','boysenberry'],
    Meats: {
        Beef: ['Filet Mignon', 'Hamburger Steak', 'Ribeye', 'Sirloin'],
        Chicken: ['Chicken Marsala', 'Fried Chicken', 'Nuggets'],
        Seafood: {
            Saltwater: ['Tuna','Swordfish','Salmon','Hake'],
            Freshwater: ['Bluegill','Crappie','Catfish','Trout','White Bass','Walleye'],
        },
    },
    Sweets: ['Bueno','Snickers','Milkyway'],
}

var foodSelectionTwo = ['beef','filet','steak']

function FindFood(food) {
    var foundFood = ByName(food)
    if (foundFood != null) {
        return 'Found your food!\n' + foundFood
    }
    if (foundFood == null) {
        foundFood = ByPart(food)
        return 'This might be your food.\n' + foundFood
    }
    return 'No food found, try again!'
}

function ByName(food) {

}

function ByPart(food) {

}

// function scan(obj) {
//     var k;
//     if (obj instanceof Object) {
//         for (k in obj){
//             if (obj.hasOwnProperty(k)){
//                 //recursive call to scan property
//                 scan( obj[k] );
//             }
//         }
//     } else {
//         //obj is not an instance of Object so obj here is a value
//     };
//
// };

console.log('Finding some food!')
console.log(FindFood('peach'))
console.log(FindFood('berry'))