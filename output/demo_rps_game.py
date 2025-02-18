def RPS_GAME():
  KEEP_PLAYING = True
  while KEEP_PLAYING == True:  
      DRAW = False
      WINNER_NAME = None
      PLAYER1_NAME = input("Player 1, what is your name? ")
      PLAYER2_NAME = input("Player 2, what is your name? ")
      print("Initiating game of rock, paper, scissors!")
      print(PLAYER1_NAME + " VERSUS " + PLAYER2_NAME)
      print(PLAYER1_NAME + " - Choose your gesture: ")
      PLAYER1_GESTURE = input("(rock, paper or scissors): ")
      print(PLAYER2_NAME + " - Choose your gesture: ")
      PLAYER2_GESTURE = input("(rock, paper or scissors): ")
      if PLAYER1_GESTURE == "rock"  and PLAYER2_GESTURE == "scissors"  or PLAYER1_GESTURE == "paper"  and PLAYER2_GESTURE == "rock"  or PLAYER1_GESTURE == "scissors"  and PLAYER2_GESTURE == "paper":
            WINNER_NAME = PLAYER1_NAME
      else:
            if PLAYER2_GESTURE == "rock"  and PLAYER1_GESTURE == "scissors"  or PLAYER2_GESTURE == "paper"  and PLAYER1_GESTURE == "rock"  or PLAYER2_GESTURE == "scissors"  and PLAYER1_GESTURE == "paper":
                    WINNER_NAME = PLAYER2_NAME
            else:
                    DRAW = True
      if DRAW == True:
            print("draw")
      else:
            print(WINNER_NAME + " wins")
      USER_RESPONSE = input("Would you like to play RPS again? (yes/no) ")
      if USER_RESPONSE == "yes":
            KEEP_PLAYING = True
      else:
            KEEP_PLAYING = False
RPS_GAME()