import argparse as argp
import csv
import os

from pandas import read_csv


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
    if in_args.UID in discord_ids:
        print(f"User found in: {csv_path[-24:]}")

# setup argparse
parser = argp.ArgumentParser(description='Search discord chat exporter CSV files to find a specific ID',
                             allow_abbrev=False)

# add arguments
parser.add_argument('UID', metavar='UID', type=int, help='Userid to locate')
parser.add_argument('csv_input', metavar='csv_input', type=str, help='The csv to search')


args = parser.parse_args()

main(args)
