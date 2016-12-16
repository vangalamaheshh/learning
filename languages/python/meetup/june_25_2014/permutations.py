import sys;

def get_all_permutations( elements ):
    if( len( elements ) <= 1 ):
        yield elements;
    else:
        for permutation in get_all_permutations( elements[1:] ):
            for i in range( len( elements ) ):
                yield permutation[:i] + elements[0:i] + permutation[i:];



def main():
    USAGE="Usage: permutations.py <comma separated values>";

    if len( sys.argv ) == 2:
        input_list = sys.argv[1].split(',');
        permutations = get_all_permutations( input_list );
        for permutation in permutations:
            print( permutation );
    else:
        print( USAGE );


if __name__ == '__main__':
    main();
