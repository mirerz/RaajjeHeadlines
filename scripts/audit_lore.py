import os
import re

# Nomenclature Manifest
TARGETS = {
    r"Endheri\s+Odi": "Endheri Odi",
    r"Endherio\s*di": "Endheri Odi",
    r"Odi\s+Endheri": "Endheri Odi",
}

def audit_manuscript(file_path):
    print(f"🔍 Auditing manuscript: {file_path}")
    count = 0
    with open(file_path, 'r', encoding='utf-8') as f:
        content = f.read()
        
    for pattern, replacement in TARGETS.items():
        matches = re.findall(pattern, content, re.IGNORECASE)
        for match in matches:
            if match != replacement:
                print(f"❌ Inconsistency found: '{match}' -> expected '{replacement}'")
                count += 1
                
    if count == 0:
        print("✅ 100% Nomenclature Consistency verified.")
    else:
        print(f"⚠️ Total inconsistencies: {count}")

if __name__ == "__main__":
    # Placeholder for the 10M token manuscript location
    manuscript_path = "assets/lore/manuscript_v1.txt"
    if os.path.exists(manuscript_path):
        audit_manuscript(manuscript_path)
    else:
        print(f"🚫 Manuscript not found at {manuscript_path}")
