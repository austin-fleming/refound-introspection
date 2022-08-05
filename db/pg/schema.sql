CREATE EXTENSION IF NOT EXISTS moddatetime SCHEMA extensions;

-----------------------------
-- ACCOUNT 
-----------------------------
CREATE TABLE IF NOT EXISTS account (
  id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
  wallet_address text NOT NULL,
  contract_address text,
  username text UNIQUE NOT NULL,
  social_links text[],
  avatar_url text,
  bio text,
  is_private boolean DEFAULT FALSE NOT NULL,
  email text,
  email_is_verified boolean DEFAULT FALSE NOT NULL,
  email_marketing_is_allowed boolean DEFAULT FALSE NOT NULL,
  phone_number text,
  phone_is_verified boolean DEFAULT FALSE NOT NULL,
  created_at timestamp with time zone DEFAULT timezone('utc'::text, now()) NOT NULL,
  updated_at timestamp with time zone DEFAULT timezone('utc'::text, now()) NOT NULL,
  nonce text NOT NULL
);

ALTER TABLE account ENABLE ROW LEVEL SECURITY;

CREATE TRIGGER handle_account_updated_at
  BEFORE UPDATE ON account
  FOR EACH ROW
  EXECUTE PROCEDURE moddatetime (updated_at);

-----------------------------
-- ACCOUNT FOLLOW RELATION
-----------------------------
CREATE TABLE IF NOT EXISTS account_follow_relation (
  follower uuid REFERENCES account (id) ON DELETE CASCADE,
  followee uuid REFERENCES account (id) ON DELETE CASCADE,
  PRIMARY KEY (follower, followee),
  created_at timestamp with time zone DEFAULT timezone('utc'::text, now()) NOT NULL
);

ALTER TABLE account_follow_relation ENABLE ROW LEVEL SECURITY;

-----------------------------
-- INITIATIVE
-----------------------------
CREATE TABLE IF NOT EXISTS initiative (
  id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
  contract_address text,
  title text NOT NULL,
  body text NOT NULL,
  cover_image text NOT NULL,
  target_date timestamp with time zone,
  target_funds bigint,
  created_at timestamp with time zone DEFAULT timezone('utc'::text, now()) NOT NULL,
  updated_at timestamp with time zone DEFAULT timezone('utc'::text, now()) NOT NULL
);

ALTER TABLE initiative ENABLE ROW LEVEL SECURITY;

CREATE TRIGGER handle_initiative_updated_at
  BEFORE UPDATE ON initiative
  FOR EACH ROW
  EXECUTE PROCEDURE moddatetime (updated_at);

-----------------------------
-- INITIATIVE VOTE
-----------------------------
CREATE TABLE IF NOT EXISTS initiative_vote (
  initiative uuid REFERENCES initiative (id) ON DELETE CASCADE,
  voter uuid REFERENCES account (id) ON DELETE CASCADE,
  vote smallint DEFAULT 0 CHECK (vote >= -1 and vote <= 1),
  PRIMARY KEY (initiative, voter),
  created_at timestamp with time zone DEFAULT timezone('utc'::text, now()) NOT NULL
);

ALTER TABLE initiative_vote ENABLE ROW LEVEL SECURITY;

-----------------------------
-- INITIATIVE MEMBER RELATION
-----------------------------
CREATE TABLE IF NOT EXISTS initiative_member_relation (
  initiative uuid REFERENCES initiative (id) ON DELETE CASCADE,
  follower uuid REFERENCES account (id) ON DELETE CASCADE,
  PRIMARY KEY (initiative, follower),
  membership_role text CHECK (membership_role in ('OWNER', 'ADMIN', 'MEMBER')),
  created_at timestamp with time zone DEFAULT timezone('utc'::text, now()) NOT NULL
);

ALTER TABLE initiative_member_relation ENABLE ROW LEVEL SECURITY;

-----------------------------
-- EVENT
-----------------------------
CREATE TABLE IF NOT EXISTS domain_event (
  id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
  event_type text NOT NULL,
  aggregate_id uuid NOT NULL,
  aggregate_type text NOT NULL,
  event_data text,
  channel text,
  created_at timestamp with time zone DEFAULT timezone('utc'::text, now()) NOT NULL
);

ALTER TABLE domain_event ENABLE ROW LEVEL SECURITY;