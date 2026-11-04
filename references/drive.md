---
title: Legacy Google Drive integration retired
description: Customer service-account folder connections are no longer supported
commands: []
---

# Legacy Google Drive integration retired

The customer service-account folder integration was retired on 2026-11-02 (AEO-1390).
Do not ask users to invite a service account or reconnect a folder. Legacy CLI list,
read and download requests return an integration-retired response after deployment.
Previously imported documents and their provenance are preserved. Separate document
imports and internal Ops Drive are not part of this retirement.
