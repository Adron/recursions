// The following is a minimal direct recursion call function.
var countdownDisplayString = ""

function Countdown(theCount, countdown)
{
    if (0 < theCount)
    {
        countdown = countdown + theCount + " "
        theCount = theCount - 1
        return Countdown(theCount, countdown)
    }
    return countdown.trim()
}

console.log("Result of direct recursion function.")
console.log(Countdown(3, countdownDisplayString))

