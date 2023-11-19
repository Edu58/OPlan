defmodule Oplan.Repo.Migrations.CreateEventComments do
  use Ecto.Migration

  def change do
    create table(:event_comments) do
      add :comment, :text
      add :user_id, references(:users, on_delete: :nothing)
      add :event_id, references(:events, on_delete: :nothing)

      timestamps()
    end

    create unique_index(:event_comments, [:user_id, :event_id, :comment])
  end
end
