# Release Notes

This directory is used to store release notes for each version.

The release notes for each version are stored in a subdirectory named after the release tag, following the format `vX.Y.Z`.

Each release directory will contain the following two files:

* `GITHUB_RELEASE.md`: The release notes formatted for GitHub Releases.
* `RELEASE_NOTES.adoc`: The release notes in AsciiDoc format.

For example, the release notes for version 1.2.3 would be structured as follows:

```text
release-notes/
└── v1.2.3/
    ├── GITHUB_RELEASE.md
    └── RELEASE_NOTES.adoc
```

When preparing release notes for a version that includes configuration model
migration updates, include:

- The introduced/updated `configModelVersion`.
- Startup consent behavior (interactive prompt vs non-interactive approval requirement).
- Backup and recovery behavior (backup file naming/location).
