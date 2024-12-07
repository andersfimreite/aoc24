#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *readfile(const char *filename);
int part1(const char *data);
int part2(const char *data);

int main(void) {
    char *filename = "input.txt";
    char *data = readfile(filename);
    if (data == NULL) {
        printf("Could not read file %s\n", filename);
        return 1;
    }

    int res1 = part1(data);
    printf("Part 1: %d\n", res1);

    int res2 = part2(data);
    printf("Part 2: %d\n", res2);

    free(data);

    return 0;
}

int part1(const char *data) {
    int count = 0;

    int cols = strchr(data, '\n') - data;
    int rows = strlen(data) / cols - 1;

    int *visit = calloc(rows * cols, sizeof(int));
    int *obst = calloc(rows * cols, sizeof(int));
    int guard_pos = 0;
    char guard_dir = '\0';

    int idx = 0;
    for (int i = 0; i < strlen(data); i++) {
        if (data[i] == '#') {
            obst[idx] = 1;
        } else if (data[i] == '^') {
            guard_pos = idx;
            guard_dir = '^';
        }
        if (data[i] != '\n') {
            idx++;
        }
    }

    while (1) {
        // count visited
        if (visit[guard_pos] == 0) {
            count++;
            visit[guard_pos] = 1;
        }

        // move
        int new_pos = 0;
        switch (guard_dir) {
            case '^':
                new_pos = guard_pos - cols;
                break;
            case 'v':
                new_pos = guard_pos + cols;
                break;
            case '<':
                new_pos = guard_pos - 1;
                break;
            case '>':
                new_pos = guard_pos + 1;
                break;
        }

        // check if outside board
        if (
            new_pos < 0 ||
            new_pos >= rows * cols ||
            (new_pos % cols == 0 && guard_pos % cols == cols - 1) ||
            (new_pos % cols == cols - 1 && guard_pos % cols == 0)
        ) {
            break;
        }

        // turn if hit obst
        if (obst[new_pos] == 1) {
            switch (guard_dir) {
                case '^':
                    guard_dir = '>';
                    break;
                case 'v':
                    guard_dir = '<';
                    break;
                case '<':
                    guard_dir = '^';
                    break;
                case '>':
                    guard_dir = 'v';
                    break;
            }
        } else {
            guard_pos = new_pos;
        }
    }

    free(visit);
    free(obst);

    return count;
}

int part2(const char *data) {
    int count = 0;

    int cols = strchr(data, '\n') - data;
    int rows = strlen(data) / cols - 1;

    int *visit = calloc(rows * cols, sizeof(int));
    int *obst = calloc(rows * cols, sizeof(int));
    int guard_pos_orig = 0;
    int guard_pos = 0;
    char guard_dir = '\0';

    int idx = 0;
    for (int i = 0; i < strlen(data); i++) {
        if (data[i] == '#') {
            obst[idx] = 1;
        } else if (data[i] == '^') {
            guard_pos = idx;
            guard_pos_orig = idx;
            guard_dir = '^';
        }
        if (data[i] != '\n') {
            idx++;
        }
    }

    while (1) {
        // mark visited
        if (visit[guard_pos] == 0) {
            visit[guard_pos] = 1;
        }

        // move
        int new_pos = 0;
        switch (guard_dir) {
            case '^':
                new_pos = guard_pos - cols;
                break;
            case 'v':
                new_pos = guard_pos + cols;
                break;
            case '<':
                new_pos = guard_pos - 1;
                break;
            case '>':
                new_pos = guard_pos + 1;
                break;
        }

        // check if outside board
        if (
            new_pos < 0 ||
            new_pos >= rows * cols ||
            (new_pos % cols == 0 && guard_pos % cols == cols - 1) ||
            (new_pos % cols == cols - 1 && guard_pos % cols == 0)
        ) {
            break;
        }

        // turn if hit obst
        if (obst[new_pos] == 1) {
            switch (guard_dir) {
                case '^':
                    guard_dir = '>';
                    break;
                case 'v':
                    guard_dir = '<';
                    break;
                case '<':
                    guard_dir = '^';
                    break;
                case '>':
                    guard_dir = 'v';
                    break;
            }
        } else {
            guard_pos = new_pos;
        }
    }

    for (int i = 0; i < rows * cols; i++) {
        if (visit[i] == 0 || i == guard_pos_orig) {
            continue;
        }

        int guard_pos = guard_pos_orig;
        int guard_dir = '^';

        char *obst_dir = calloc(rows * cols, sizeof(char));

        while (1) {
            // move
            int new_pos = 0;
            switch (guard_dir) {
                case '^':
                    new_pos = guard_pos - cols;
                    break;
                case 'v':
                    new_pos = guard_pos + cols;
                    break;
                case '<':
                    new_pos = guard_pos - 1;
                    break;
                case '>':
                    new_pos = guard_pos + 1;
                    break;
            }

            // Check if outside board
            if (
                new_pos < 0 ||
                new_pos >= rows * cols ||
                (new_pos % cols == 0 && guard_pos % cols == cols - 1) ||
                (new_pos % cols == cols - 1 && guard_pos % cols == 0)
            ) {
                break;
            }


            if (obst[new_pos] == 1 || new_pos == i) {
                // if hitting obst from same dir as before, must be looping
                if (obst_dir[new_pos] == guard_dir) {
                    count++;
                    break;
                }

                obst_dir[new_pos] = guard_dir;

                switch (guard_dir) {
                    case '^':
                        guard_dir = '>';
                        break;
                    case 'v':
                        guard_dir = '<';
                        break;
                    case '<':
                        guard_dir = '^';
                        break;
                    case '>':
                        guard_dir = 'v';
                        break;
                }

            } else {
                guard_pos = new_pos;
            }
        }

        free(obst_dir);
    }

    free(visit);
    free(obst);

    return count;
}

char *readfile(const char *filename) {
    FILE *f = fopen(filename, "rb");
    if (f == NULL) {
        return NULL;
    }

    fseek(f, 0, SEEK_END);
    long filesize = ftell(f);
    fseek(f, 0, SEEK_SET);

    char *buf = malloc(filesize + 1);
    if (buf == NULL) {
        fclose(f);
        return NULL;
    }

    fread(buf, sizeof(char), filesize, f);
    buf[filesize] = '\0';

    fclose(f);

    return buf;
}
