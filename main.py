import argparse as argp
import csv
import os

from pandas import read_csv

unique_ids = []
unique_old = []


def main(in_args):
    # reading CSV file
    user_data = read_csv(in_args.csv_input)
    discord_ids = user_data['AuthorID'].tolist()

    if os.path.isfile(in_args.csv_output) is False:
        write_csv(["User ID"], in_args.csv_output)
    else:
        old_data = read_csv(in_args.csv_output)
        for x in old_data['User ID'].tolist():
            unique_old.append(int(x))

    for user_id in discord_ids:
        int_id = int(user_id)
        print("User: ", str(int_id), end="\r", flush=True)
        if user_id not in unique_ids and user_id not in unique_old:
            unique_ids.append(user_id)

    for user_id in unique_ids:
        print("User: ", str(user_id), end="\r", flush=True)
        update_csv(user_id, in_args.csv_output)


def write_csv(fields, csv_path):
    with open(csv_path, 'w', newline='') as csv_file:
        writer = csv.DictWriter(csv_file, fieldnames=fields)

        writer.writeheader()


def update_csv(uids, csv_path):
    # print("FPS: ", str(FPS), end="\r", flush=True)
    with open(csv_path, 'a', newline='') as csv_file:
        writer = csv.writer(csv_file)
        writer.writerow([uids])


# setup argparse
parser = argp.ArgumentParser(description='Make a list of all unique user IDs from a discord chat exporter CSV',
                             allow_abbrev=False)

# add arguments
parser.add_argument('csv_input', metavar='csv_input', type=str, help='The path of your exported CSV file')
parser.add_argument('csv_output', metavar='csv_output', type=str, help='The path your new CSV with unique IDs')

args = parser.parse_args()

main(args)
