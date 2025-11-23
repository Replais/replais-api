-- Remove ONLY system personas that were inserted in the UP migration
DELETE FROM personas
WHERE user_id IS NULL
  AND slug IN (
    'default',
    'buddy',
    'ceo',
    'empath',
    'charm',
    'direct'
  );