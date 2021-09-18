/**
 * NCurses test to check creature
 */
#include<stdio.h>
#include<stdlib.h>

#include<ncurses.h>


// Definitions
void creature();

int main(int argc, char *argv[]) {
    initscr();
    keypad(stdscr, TRUE);
    noecho();

    /* Check if has colors */
    if (has_colors() == FALSE) {
        endwin();
        puts("Your terminal does not support color");
        exit(1);
    }

    printw("Hello");
    creature();
    refresh();
    getch();
    endwin();
}


void creature() {
    start_color();
    // Body
    init_pair(1, COLOR_BLUE, COLOR_BLUE);
    // eyes
    init_pair(2, COLOR_WHITE, COLOR_WHITE);
    // terminators
    init_pair(3, COLOR_GREEN, COLOR_GREEN);

    attron(COLOR_PAIR(1));
    mvprintw(2, 13, "######");
    mvprintw(3, 10, "####");
    mvprintw(3, 18, "####");
    mvprintw(4, 8, "#########");
    mvprintw(5, 5, "####");
    mvprintw(5, 10, "####");
    mvprintw(5, 15, "####");
    mvprintw(6, 2, "######");
    mvprintw(6, 11, "##");
    mvprintw(6, 16, "######");
    mvprintw(7, 0, "#########");
    mvprintw(7, 10, "####");
    mvprintw(7, 15, "#########");
    mvprintw(8, 0, "####");
    mvprintw(8, 6, "##############");
    mvprintw(8, 21, "###");
    mvprintw(9, 0, "####");
    mvprintw(9, 10, "#########");
    mvprintw(10, 0, "########################");
    mvprintw(11, 2, "####################");
    mvprintw(12, 5, "##############");
    attroff(COLOR_PAIR(1));

    attron(COLOR_PAIR(2));
    mvprintw(5, 9, "+");
    mvprintw(5, 14, "+");
    mvprintw(6, 8, "+++");
    mvprintw(6, 13, "+++");
    mvprintw(7, 9, "+");
    mvprintw(7, 14, "+");
    attroff(COLOR_PAIR(2));

    attron(COLOR_PAIR(3));
    mvprintw(8, 4, "**");
    mvprintw(8, 19, "**");
    mvprintw(9, 4, "*******");
    mvprintw(9, 19, "********");
    mvprintw(13, 8, "***");
    mvprintw(13, 13, "***");
    mvprintw(14, 8, "***");
    mvprintw(14, 13, "***");
    attroff(COLOR_PAIR(3));

}


//              #///(%
//           %((%    #/#&
//         #//////#&
//      %(//*////*//(%
//   &#////***//***////#&
// %(///////*////*///////(%
// %(/(#(////////////(#(/(%
// %(/(######(///////(######%
// %(////////////////////(%
//   &#////////////////#&
//      %(//////////(%
//         %#%  %#%
//         %#%  %#%

