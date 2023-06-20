func convertTemperature(celsius float64) []float64 {
    var result []float64
    
    kelvin := celsius + 273.15
    fahrenheit := celsius * 1.8 + 32
    
    result = append(result, kelvin, fahrenheit)
    
    return result
}