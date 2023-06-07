package hbery_aoc2015

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

type MoleculeReplacement struct {
	from	string
	to		string
}

func (mr *MoleculeReplacement) to_string() string {
	return fmt.Sprintf("ModuleReplacement{from: %s, to: %s}", mr.from, mr.to)
}

func parse_molecule_replacemts(input string) ([]*MoleculeReplacement, string) {
	var mreps []*MoleculeReplacement
	var molecule string


	lines := strings.Split(input, "\n")

	for i, line := range lines {
		if len(line) == 0 { continue }
		if i == len(lines) - 1 {
			molecule = line
			break
		}

		el := strings.Split(line, " => ")
		mreps = append(mreps, &MoleculeReplacement{
			from: el[0],
			to:   el[1],
		})
	}

	return mreps, molecule
}

func remove_molecule_duplicates(input *[]string) *[]string {
    distinct_strings := make(map[string]bool)
    deduped_slice := []string{}
    for _, item := range *input {
        if _, value := distinct_strings[item]; !value {
            distinct_strings[item] = true
            deduped_slice = append(deduped_slice, item)
        }
    }
    return &deduped_slice
}

func create_every_molecule_replacement(molecule string, mreps *[]*MoleculeReplacement) *[]string {
	var created_molecules []string = []string{}

	for _, mr := range *mreps {
		mr_count := strings.Count(molecule, mr.from)
		mr_idx := strings.Index(molecule, mr.from)
		mr_lidx := strings.LastIndex(molecule, mr.from)
		mr_split_mol := strings.Split(molecule, mr.from)
		var new_molecule string
		var start, end int

		if mr_count == 0 { continue }

		if mr_idx == 0 && mr_lidx == len(molecule) - len(mr.from) {
			start = 1
			end = len(mr_split_mol) - 1
		} else if mr_idx == 0 {
			start = 0
			end = mr_count
		} else if mr_lidx == len(molecule) - len(mr.from) {
			start = 1
			end = len(mr_split_mol) - 1
		} else { start = 1
			end = mr_count
		}

		for start <= end {
			new_molecule = fmt.Sprintf("%s%s%s", strings.Join(mr_split_mol[:start], mr.from), mr.to, strings.Join(mr_split_mol[start:], mr.from))

			created_molecules = append(created_molecules, new_molecule)
			start++
		}
	}

	return &created_molecules
}

func fabricate_molecule(molecule string, entry_replacement *MoleculeReplacement, molecule_replacements *[]*MoleculeReplacement) int64 {
	var op_count int64 = 0

	// combinate
	for molecule != entry_replacement.from {
		if len(molecule) == 1 && molecule == entry_replacement.to {
			op_count++
			molecule = entry_replacement.from
		}

		rand_mr := (*molecule_replacements)[rand.Intn(len(*molecule_replacements))]

		if strings.Contains(molecule, rand_mr.to) {
			molecule = strings.Replace(molecule, rand_mr.to, rand_mr.from, 1)
			op_count++
		}
	}

	return op_count
}

func day19_p1(input string) (int64, error) {
	var result int64 = 0
	var molecule_replacements []*MoleculeReplacement
	var molecule string
	
	molecule_replacements, molecule = parse_molecule_replacemts(strings.TrimSuffix(input, "\n"))

	created_molecules := create_every_molecule_replacement(molecule, &molecule_replacements)
	created_molecules = remove_molecule_duplicates(created_molecules)
	result = int64(len(*created_molecules))

	return result, nil
}

func day19_p2(input string) (int64, error) {
	var result int64 = 0
	var molecule_replacements []*MoleculeReplacement
	var entry_replacement *MoleculeReplacement = &MoleculeReplacement{from: "e", to: "O"}
	var molecule string

	molecule_replacements, molecule = parse_molecule_replacemts(strings.TrimSuffix(input, "\n"))

	result = fabricate_molecule(molecule, entry_replacement, &molecule_replacements)

	return result, nil
}

func Solution_Day19(part int, input string) (int64, error) {
	if part == 1 {
		return day19_p1(input)
	} else if part == 2 {
		return day19_p2(input)
	} else {
		return 0, errors.New(fmt.Sprintf("Error: Hold on cowboy. No such part (%d) of this day task", part))
	}
}
