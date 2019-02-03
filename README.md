# resman
Resman is a resource manager. It may become a polivalent tool able to manage many different type of resources, but for now it just handles local git repositories.

# Intended usage
`resman git [--root path] [--exclude path]`: print a list of local git repos with their current status

# To Do
* List first my personal repos, then my organizations', then write access, last others
* Implement resman git --owner (user|org) --host (local|github|bitbucket)
* Implement describe repo?
* Color repos based on host
* Option to show size and color it like exa
* Show git info in tabs, and customize what tabs to show
* Detect if master is not synced with origin
