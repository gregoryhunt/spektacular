## Step {{step}}: {{title}}

plan.md has been written to `{{plan_path}}`.

Now pipe the filled context.md the same way — spektacular will write it to `{{context_path}}`:

```
cat <<'EOF' | {{config.command}} plan goto --data '{"step":"{{next_step}}"}' --stdin context_template
<complete filled context.md here>
EOF
```
