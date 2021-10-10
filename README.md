# Longform

Longform is a server that allows for storing, editing, and hosting academic/scholarly work.

This project is not ready for use yet. Here is a list of desired future functionality:

* document page
    * easy linking of headers and footnotes
    * readability on all screen sizes
    * light and dark themes
* document editing page
    * main body of document. Can be simple markdown for now, but should allow linking to footnotes and glossary terms somehow
    * footnotes - these are per-document
    * glossary terms - these are global (my use-case for now involves a single theme for all published documents), though I may want to consider having an option for per-document glossary terms in the future
* list page
    * titles and blurbs for each page
    * should have some sort of subdirectory/recursive listing functionality
* list edit page
    * complete freedom to reorganize/reorder documents and how they will be published

For now, I plan on using a sqlite3 db for data that needs to be preserved, and the `static` directory can be used to publish any static files.

The UI will start out very basic. I want to first figure out how the data and application will interact. Then, I can focus on improving the experience.
