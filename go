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
    cout<< "Welcome" + first_name + "! Try your best to escape the federation" <<endl;
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
                cout<< "Hi friend! My name is Pomme! I'm stuck in here will you help me? I need to find my brother Dapper! He likes dark places and cults! Yes (L) No (R)"<<endl;
                cin>> position;
                if(position == 'L'){
                    cout<< "Thank you so much! She grabs your hand and leads you to a door behind her. She opens the door and steps outside, revealing another long hallway." << endl;
                    cout<< "Should we go left (L) or right(R)?"<<endl;
                    cin>>position;
                    if(position == 'L'){
                        cout<< "As you walk down the hallway, you fall into a pit. Pomme sadly waves goodbye as you fall to your death."<< endl;
                        cout<< "Start over? Yes(L) No (R)"<<endl;
                        cin>> position;
                    }
                    if(position == 'R'){
                        cout << "Follow me pute! She grabs your hand and pulls you down the hallway."<< endl;
                        cout << "Should we go follow the red trail(L) on the ground or follow the noise(R)" << endl;
                        cin>> position;
                        if(position == 'L'){
                            cout<< "I see him! As you round the bend, you see a boy that looks similar to Pomme in the middle of a summoning circle."<<endl;
                            cout<< "Will you, run in and stop him (L) watch in silence (R)" << endl;
                            cin>> position;
                            if(position == 'L'){
                                cout<<"The second you take a step into the summoning circle you explode into a red vapor"<< endl;
                                cout<< "Start over? Yes(L) No (R)"<<endl;
                                cin>> position;
                            }
                            if(position == 'R'){
                                cout<< "As Dapper finishes his chanting you're suddenly teleport outside" <<endl;
                                cout<< "You have escaped! Congrats!" << endl;
                                cout<< "Start over? Yes(L) No (R)"<<endl;
                                cin>> position;
                            }

                        }
                        if(position == 'R'){
                            cout<<"As you round the bend, you see a large monster with thousands of teeth. Pomme looks up. Um thats not Dapper." << endl;
                            cout<<"Pomme runs away as you're eaten." <<endl;
                            cout<< "Start over? Yes(L) No (R)"<<endl;
                            cin>> position;
                        }
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
            cout << "As you walk down the hallway, a child with bright blonde hair runs straight into you."<<endl;
            cout <<"Hi! My name is Chayanne! I'm looking my sister, her name is Tallulah. Can you help? Yes(L) No(R)" <<endl;
            cin>>position;
            if(position == 'L'){
                cout<< "Thank you so much! He grabs your hand and continues sprinting down the hallway." <<endl;
                cout<<"She likes music and gardening! I hope shes ok." <<endl;
                cout<< "Do you, follow the noise (L) follow the sunlight (R)" <<endl;
                cin>> position;
                if(position == 'L'){
                    cout<< "As you walk towards the noise, a piano falls on your head." <<endl;
                    cout<< "Start over? Yes(L) No (R)"<<endl;
                    cin>> position;
                }
                if(position == 'R'){
                    cout<< "I see her! Chayanne runs towards a girl looking at some flowers." <<endl;
                    cout<< "Now that I've found Tallulah I can lead us out!"<<endl;
                    cout<< "Do you follow him(L) or find your own way (R)"<<endl;
                    cin>> position;
                    if(position == 'L'){
                        cout<< "Chayanne confidently leads the way outside." <<endl;
                        cout<< "You have escaped! Congrats!" << endl;
                        cout<< "Start over? Yes(L) No (R)"<<endl;
                        cin>> position;
                    }
                    if(position == 'R'){
                        cout<< "You are lost forever in the maze." <<endl;
                        cout<< "Start over? Yes(L) No (R)"<<endl;
                        cin>> position;
                    }
                }
            }
            if(position == 'R'){
                cout<< "Chayanne pulls out a giant sword and skewers you with it." <<endl;
                cout<< "Start over? Yes(L) No (R)"<<endl;
                cin>> position;
            }
        }
        //return 0;
}
        
}
        

