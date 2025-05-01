#include <iostream>
#include <iomanip>
#include <string>
using namespace std;

int main()
{
    //Declare variables
    string first_name, last_name;
    char position;

    //Capture user input
    cout << "What is your name?" << endl;
    cin>> first_name;
    cout << "Which way will you go first? Enter L for left, R for right:" << endl;
    cin >> position;

    //Validate input
    
    string message = composeMessage(first_name);
    cout << message << endl;
    return 0;
}

    //Compose and display message
    string composeMessage(string first_name, string last_name, char position){
        string what_role;

        if(position == 'L'){
        what_role = "The passenger rooms are thru the third door on the right";
        //return 0;
}
        if(position == 'R') {
        what_role = "The captain's quarters are up the stairs on the left";
       // return 0;
}
        
        

        string message = "Welcome Aboard " + first_name + " " + last_name + "! " + what_role;
        return message;
    }
        

