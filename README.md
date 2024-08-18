### Tool Description

The `notestool` is a command-line application designed to manage collections of short single-line notes. Each collection is stored in a separate database file, where notes can be added, viewed, or removed.

### Usage

#### Command Syntax

```bash
$ ./notestool <collection_name>
```

#### Examples

1.  **Create or manage a collection:**
    
    ```bash
    $ ./notestool coding_ideas
    ```
    
    This command will either create a new collection named `coding_ideas` or load an existing one if it already exists.
    
2.  **Displaying notes:**
    
    After loading a collection, the tool displays a menu where you can choose to:
    
    *   Display all notes
    *   Add a new note
    *   Remove an existing note
    *   Exit the program
3.  **Help message:**
    
    If no argument is provided, more than one argument is given, or the argument is `help`, the tool displays a brief help message:
    
    ```bash
    $ ./notestool help
    Usage: ./notestool [COLLECTION]
    ```
    

### Data Storage

*   Each collection is stored as a plain text file with the same name as the collection (e.g., `coding_ideas.txt`).
*   Notes within each collection are stored as separate lines in the text file.