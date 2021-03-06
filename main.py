import argparse as argp
import csv
import os

from pandas import read_csv

unique_ids = []
unique_old = []


def main(in_args):

    if not os.path.isfile(in_args.csv_input):
        for filename in os.listdir(in_args.csv_input):
            file_path = os.path.join(in_args.csv_input, filename)
            process(in_args, str(file_path))
    else:
        # reading CSV file
        csv_path = in_args.csv_input
        process(in_args, csv_path)

    print("Done.")


def process(in_args, csv_path: str):
    user_data = read_csv(csv_path)
    discord_ids = user_data['AuthorID'].tolist()
    # check if the output exists to add new IDs too
    if os.path.isfile(in_args.csv_output) is False:
        write_csv(["User ID"], in_args.csv_output)
    else:
        old_data = read_csv(in_args.csv_output)
        for x in old_data['User ID'].tolist():
            unique_old.append(int(x))
    # check for new IDs
    for user_id in discord_ids:
        int_id = int(user_id)
        print("User: ", str(int_id), end="\r", flush=True)
        if user_id not in unique_ids and user_id not in unique_old:
            unique_ids.append(user_id)
    # add users to the CSV
    for user_id in unique_ids:
        print("User: ", str(user_id), end="\r", flush=True)
        update_csv(user_id, in_args.csv_output)

    unique_ids.clear()


def write_csv(fields, csv_path):
    """Create initial csv file"""
    with open(csv_path, 'w', newline='') as csv_file:
        writer = csv.DictWriter(csv_file, fieldnames=fields)

        writer.writeheader()


def update_csv(uids, csv_path):
    """Add ids into CSV"""
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
