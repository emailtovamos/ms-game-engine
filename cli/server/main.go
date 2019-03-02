package main
// Connect to storage serve to get latest 4 scores
// If the trend is worse bestscore then return larger size of shape to make the game less difficult
// If the trend is improved bestscore then return smaller size of shape to make the game more difficult
// Start with just assuming you have the scores from storage by randomly generating 4 numbers with certain order
// TEST if the function returns values as expected
// THEN add logic to connect to storage service once storage service is up and running

// ms-game-engine opens a server so that when it gets a GetSize request, it returns size of shape: height, width
// Only border-radius, height and width matter 
// Border radius (50% - circle, 0% - rectangle) may or may not be decided on the client side
func main() {

}