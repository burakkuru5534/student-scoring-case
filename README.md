# student-scoring-case

# Introduction

In this project we aimed to provide a solution to the problem of scoring students.
Rules:
   1. A student can get multiple points in one day
   2. The teacher can give points at any time of the day
   3. The lowest possible score: -5
   4. The highest possible score: +5
   5. Students are divided into two groups; A and B groups
   6. The first 10 students in group A with the highest scores and the next 5 students in group B are ranked according to their scores.
   7. When the score of the student in the first row of group B is higher than the score of the student in the last row of group A, the students change places.
   8. The lists are updated every hour and extra points are awarded to students in group A every hour;
       1. 10, 9, 8, 7, students are given 1
       2. 6, 5, 4, 3, Students are given 2
       3. 2 Students in the 1st place are given 3 points
       4. Lists are reset at the beginning of each week
   9. Student numbers range from 101 to 115
   10. Everyone has zero points at the beginning

# Installation

Golang must be installed on your computer. You can download it from [here](https://golang.org/dl/).
Postgresql must be installed on your computer. You can download it from [here](https://www.postgresql.org/download/).

you should initialize your database with the following command:

```bash
$ psql -U postgres -f db_structure.sql
```

go build -o student-scoring-case

# Usage

```bash
$ ./student-scoring-case
```


# Solution

a ticker had been created that runs every hour and updates the scores of students and the ranking list.
1 hour assumed as 10 seconds for testing purposes.
user (teacher) can use the terminal to add points to students.
the teacher can add points to students at any time of the day.
the teacher can add points to students in any range from -5 to +5.
the teacher can add points to students in any group (A or B).
the teacher can add points to students in any number from 101 to 115.
the teacher can end the day by selecting the option 1 in the menu.
the teacher can give points by selecting the option 2 in the menu.
after that selection the teacher can select the group of the student.
after that selection the teacher can enter the points that he/she wants to give to the student.
the teacher can see the ranking list by selecting the option 3 in the menu.
the teacher can end the week by selecting the option 4 in the menu.
the teacher can end the program by selecting the option 5 in the menu.

# Conclusion

we have created a program that can score students and rank them according to their scores.
we have used the following technologies:
1. Golang
2. Postgresql


