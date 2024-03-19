# JukeBox application

## Problem

JukeBox application is intended to maintain a catalogue of music albums by various musicians. The application should also maintain information about musicians, music and the relationships between musicians and music (i.e, who has sung/played which album). You are required to implement a set of demonstrable REST APIs that can be consumed by any UI/client application to retrieve information about artists and music albums & also to add new music albums, artist details.

Following information should be maintained for Music Albums

- Album name (Mandatory, Should be minimum 5 characters)
- Date of release (Mandatory)
- Genre (i.e., Type of music)
- Price (Mandatory, value between 100 to 1000)
- Description

Following information should be maintained for Musicians

- Name (Mandatory, Should be minimum 3 characters)
- Musician type (Vocalist, Instrumentalist, etc...)

Constraints

- A music album can be sung/played by one or more musicians
- A musician can contribute (sing or play) to multiple music albums

APIS to be Developed

Implement REST APIs for the following:

1. API to create/update music album records
2. API to create/update musician records
3. API to retrieve the list of Music albums sorted by Date of release in ascending order (i.e
Oldest first)
4. API to retrieve the list of Music albums for a specified musician sorted by Price in ascending order (i.e Lowest first)
5. API to retrieve the list of musicians for a specified music album sorted by musician's Name in ascending order.

### Evaluation Criteria:
- Completeness of Data model
- Number of APIs completed.
- Problem solving ability
- Code Quality
- Adherence to coding standards
- Implementation of Unit test cases and code coverage
- Effective use of Git repository, i.e., Task breakdown and Issue creation, Separate commits per task against issue.