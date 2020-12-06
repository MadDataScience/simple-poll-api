CREATE TABLE IF NOT EXISTS Polls(
    PollID INTEGER PRIMARY KEY AUTOINCREMENT,
    UserID INT ,
    Title VARCHAR (255) ,
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
)
