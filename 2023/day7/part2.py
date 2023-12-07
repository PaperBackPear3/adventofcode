from sys import argv
import os
import math

file_name = argv[1]

dir_path = os.path.dirname(os.path.realpath(__file__))
file1 = open(dir_path + "/" + file_name, "r")
read_lines = file1.read().splitlines(keepends=False)


map_letters_to_numbers = {
    "A": 14,
    "K": 13,
    "Q": 12,
    "J": 1,
    "T": 10,
}


def get_hand_type(cards):
    # 7 possible cases
    cards_copy = cards
    # J now counts the best possible hand
    val = 1
    if cards_copy == [1, 1, 1, 1, 1]:
        return 7
    cards_occurencies_dict = {i: cards_copy.count(i) for i in cards_copy}
    # check if theres a J (1) in the hand and make it count as the most present one
    if 1 in cards_occurencies_dict.keys():
        qty = cards_occurencies_dict.pop(1)
        cards_occurencies_dict[
            max(cards_occurencies_dict, key=cards_occurencies_dict.get)
        ] += qty
    match len(cards_occurencies_dict):
        case 1:  # all cards are the same
            val = 7
        case 2:  # can be four of a kind or full
            for key in cards_occurencies_dict:
                if cards_occurencies_dict[key] == 4:
                    val = 6
                elif cards_occurencies_dict[key] == 3:
                    val = 5
        case 3:  # can be three of a kind or two pairs
            for key in cards_occurencies_dict:
                if cards_occurencies_dict[key] == 3:
                    val = 4
                elif cards_occurencies_dict[key] == 2:
                    val = 3
        case 4:  # can be pair
            val = 2

    return val


def get_highest_hand(first_hand, second_hand):
    # return the highest hand between two hands when they are the same type, who has the higher cards wins
    for index, card in enumerate(first_hand):
        if card > second_hand[index]:
            return first_hand
        elif card < second_hand[index]:
            return second_hand


def prepare_input(lines):
    # availablw acrd from strongest to weakest
    # A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, 2

    players = []
    for line in lines:
        [cards_in_hand, bets] = line.split(" ")
        bets = int(bets)
        cards_as_numbers = []
        for card in cards_in_hand:
            if card in map_letters_to_numbers:
                cards_as_numbers.append(map_letters_to_numbers[card])
            else:
                cards_as_numbers.append(int(card))
        players.append([cards_as_numbers, bets])
    return players


def get_player_rank_in_game(index, all_players):
    # return the rank of the player in the game
    playing_player = all_players[index]
    player_rank_score = 1
    playing_player_hand_type = get_hand_type(playing_player[0])
    for inner_index, player in enumerate(all_players):
        if inner_index == index:
            continue
        player_hand_type = get_hand_type(player[0])
        if playing_player_hand_type > player_hand_type:
            player_rank_score += 1
        elif playing_player_hand_type == player_hand_type:
            if get_highest_hand(playing_player[0], player[0]) == playing_player[0]:
                player_rank_score += 1
    return player_rank_score


all_players = prepare_input(read_lines)

res = 0
for index, player in enumerate(all_players):
    rank = get_player_rank_in_game(index, all_players)
    res += rank * player[1]

print(res)
