package main
import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    objScanner := bufio.NewScanner(os.Stdin)
    
    objScanner.Scan()
    
    intSearchValue,_ := strconv.Atoi(objScanner.Text())
    
    objScanner.Scan()
    
    intAryLength,_ := strconv.Atoi(objScanner.Text())
    
    objScanner.Split(bufio.ScanWords)
    
    var aryData = make([]int, intAryLength)
    
    for intCount:=0; intCount< intAryLength; intCount++ {
        objScanner.Scan()
        intTemp,_ :=  strconv.Atoi(objScanner.Text())
                
        aryData[intCount] = intTemp
    }
        
    fmt.Printf("%d", binarySearch(aryData[:],intSearchValue))
}

func binarySearch(pAryValues []int, pIntValue int) int {
    var intMaxIndex int = len(pAryValues) - 1
    
    var intStartIndex int = 0
    
    var intMidIndex int = intMaxIndex / 2
    
    for intStartIndex <= intMaxIndex {
        
        if(pAryValues[intMidIndex] == pIntValue){
            return intMidIndex
        }else if pAryValues[intMidIndex] > pIntValue {
            intMaxIndex = intMidIndex -1
        }else{
            intStartIndex = intMidIndex + 1
        }
        
        intMidIndex = (intMaxIndex + intStartIndex) /2
    }
    
    return -1
} 
