defmodule Oplan.Repo.Migrations.AlterTableUsersAddNamesAndAccountType do
  use Ecto.Migration

  def change do
    create table(:account_types) do
      add :name, :string, default: "normal"

      timestamps()
    end

    create unique_index(:account_types, [:name])

    alter table(:users) do
      add :first_name, :string
      add :last_name, :string
      add :username, :string
      add :account_type, references(:account_types, on_delete: :delete_all), null: false
    end

    create unique_index(:users, [:username])

    create table(:profiles) do
      add :user_id, references(:users, on_delete: :delete_all), null: false
      add :avatar, :string
      add :date_of_birth, :string
      add :phone_number, :string
      add :nationality, :string
      add :city, :string

      timestamps()
    end

    create unique_index(:profiles, [:phone_number])
  end
end
