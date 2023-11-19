defmodule Oplan.Repo.Migrations.CreateEventFollowers do
  use Ecto.Migration

  def change do
    create table(:event_followers) do
      add :user_id, references(:users, on_delete: :nothing)
      add :event_id, references(:events, on_delete: :nothing)

      timestamps()
    end

    create index(:event_followers, [:user_id])
    create index(:event_followers, [:event_id])
  end
end
