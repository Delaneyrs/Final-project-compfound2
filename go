#include <iostream>
#include <iomanip>
#include <string>
using namespace std;

int main()
{
    //Declare variables
    string first_name, last_name;
    char position;
    string journey;

    //Capture user input
    cout << "What is your name?" << endl;
    cin>> first_name;
    cout<< "Welcome" + first_name + "!" <<endl;
    cout << "Which way will you go first? Enter L for left, R for right:" << endl;
    cin >> position;


        if(position == 'L'){
        cout<< first_name+ " goes to the left. As they walk down the dark hallway, they see a room." << endl;
        cout<<"Do you, go through the door? (L) Or continue down the hallway?(R)" << endl;
        cin>> position;
        if(position == 'L'){
            cout<< "As you walk through the door, it shuts behind you. The room is shrouded in complete darkness, but you hear a voice from the far side of the room. "<<endl;
            cout<<"Do you, walk towards the noise? (L) Or try to leave through the closer door? (R)" << endl;
            cin>> position;
            if(position == 'L'){
                cout<<"As " + first_name + " starts walking towards the noise, they suddenly bump into something."<< endl;
                cout<< "Hi friend! My name is Pomme! I'm stuck in here will you help me? Yes (L) No (R)"<<endl;
                cin>> position;
                if(position == 'L'){
                    cout<< "Thank you so much! She grabs your hand and leads you to a door behind her. She opens the door and steps outside, revealing another long hallway." << endl;
                    cout<< "Should we go left (L) or right(R)?"<<endl;
                    cin>>position;
                    if(position == 'L'){
                        
                    }
                    if(position == 'R'){

                    }
                }
                if(position == 'R'){
                    cout<<"She looks at you with tears in her eyes before starting to cry."<< endl;
                    cout<< "She pulls out a comically large hammer and crushes you." <<endl;
                    cout<< "Start over? Yes(L) No (R)"<<endl;
                    cin>> position;
                }
            }
            if(position == 'R'){
                cout<< "You try to open the door, but does not budge. You are there for all of eternity with your new friend."<<endl;
                cout<< "Start over? Yes(L) No (R)"<<endl;
                cin>> position;
            }
        }
        if(position == 'R'){

        }
        //return 0;
}
        if(position == 'R') {
        journey = "The captain's quarters are up the stairs on the left";
       // return 0;
}
        
}
        

