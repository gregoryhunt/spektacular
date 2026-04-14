## Step {{step}}: {{title}}

context.md has been written to `{{context_path}}`.

Now pipe the filled research.md — spektacular will write it to `{{research_path}}`:

```
cat <<'EOF' | {{config.command}} plan goto --data '{"step":"{{next_step}}"}' --stdin research_template
<complete filled research.md here>
EOF
```
