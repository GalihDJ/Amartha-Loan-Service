CREATE TABLE "investor_list" (
  "investor_id" varchar,
  "investor_name" varchar,
  "investor_email" varchar,
  "created_date" timestamptz,
  PRIMARY KEY ("investor_id")
);

INSERT INTO investor_list (investor_id,investor_name,investor_email,created_date) VALUES
	 ('276883cb-e718-4238-ae9e-13a064503efa','Investor Alpha','investor_a@gmail.com','2025-03-04 17:15:06.163419+07'),
	 ('784cffc1-2473-47dc-93f9-b566c8960a4a','Investor Beta','investor_b@gmail.com','2025-03-05 05:43:49.446982+07'),
	 ('4d49edeb-87e8-4579-93b8-737e36db494b','Inevstor Gamma','investor_g@gmail.com','2025-03-05 05:44:07.93488+07');


CREATE TABLE "loan_request" (
  "loan_request_id" varchar,
  "borrower_id" varchar,
  "principal_amount" float,
  "rate" float,
  "roi" float,
  "state" varchar,
  "created_date" timestamptz,
  PRIMARY KEY ("loan_request_id")
);


CREATE TABLE "loan_approval" (
  "loan_approval_id" varchar,
  "loan_request_id" varchar,
  "field_validator_proof" varchar,
  "employee_id" varchar,
  "approved_date" timestamptz,
  PRIMARY KEY ("loan_approval_id")
);

CREATE TABLE "investment_list" (
  "investment_id" varchar,
  "loan_request_id" varchar,
  "investor_id" varchar,
  "amount" float,
  "created_date" timestamptz,
  PRIMARY KEY ("investment_id")
);

CREATE TABLE "loan_disbursement" (
  "disbursement_id" varchar,
  "loan_request_id" varchar,
  "agreement_letter_url" varchar,
  "employee_id" varchar,
  "disbursement_date" timestamptz,
  PRIMARY KEY ("disbursement_id")
);

CREATE TABLE "user_list" (
  "user_id" varchar,
  "name" varchar,
  "email" varchar,
  "user_role" varchar,
  "created_date" timestamptz,
  PRIMARY KEY ("user_id")
);

INSERT INTO user_list (user_id,"name",email,user_role,created_date) VALUES
	 ('EMP-090AX71','Jane Doe','janedoe@email.com','EMPLOYEE','2025-03-05 12:38:17.339+07'),
	 ('USR-210GH80','John Doe','johndoe@email.com','USER','2025-03-05 12:38:17.339+07');




